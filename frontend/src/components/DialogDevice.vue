<template>
  <v-card-text>
    <v-form ref="form" v-if="!created">
      <v-row>
        <v-col cols="8"><v-text-field v-model="payload.name" label="Name" :rules="[(val: string) => validationRules.required(val, 'name')]" /></v-col>
        <v-col cols="4"><v-select :items="['laptop', 'mobile']" v-model="payload.type" label="Device Type" /></v-col>
      </v-row>
      <v-row>
        <v-col cols="12"><v-text-field v-model="payload.description" label="Notes"></v-text-field></v-col>
      </v-row>
      <v-row v-if="!sameUser">
        <v-col cols="12"><v-select v-model="payload.user" label="User" :items="existingUsers" autocomplete @input="fetchUsers" /></v-col>
      </v-row>
      <v-expansion-panels class="pt-4 elevated-0">
        <v-expansion-panel title="Advanced Settings">
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="9"><v-text-field v-model="payload.endpoint" placeholder="10.254.0.1" label="Endpoint" hide-details /></v-col>
              <v-col cols="3"><v-checkbox v-model="payload.keepAlive" label="Persistent Keep Alive" hide-details /></v-col>
            </v-row>
            <v-row>
              <v-col><v-text-field v-model="payload.allowedIps" label="Allowed IPs" hide-details /></v-col>
              <v-col><v-text-field v-model="payload.dns" label="DNS servers" hide-details /></v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-form>
    <template v-else>
      <code class="block whitespace-pre overflow-x-scroll"><pre>{{ getPeerConfig(peerConfig) }}</pre></code>
    </template>
  </v-card-text>
</template>

<script setup lang="ts">
import { generateKeyPair } from '@stablelib/x25519';
import { encode as b64encode } from '@stablelib/base64';
import type { CreateDeviceRequest } from '@/lib/api';
import { API_URL, validationRules, type FilterArgs } from '@/lib/utils';
import { onMounted, ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useDataStore } from '@/stores/data';
import { ValidationError } from '@/lib/errors';
import { useDataDialogStore } from '@/stores/dialogs';
import { useIdentitiesStore } from '@/stores/identities';

const identities = useIdentitiesStore();
const dialogStore = useDataDialogStore();
const authStore = useAuthStore();
const created = ref(false);
const form = ref();

const props = defineProps({
  sameUser: {
    type: Boolean,
  }
})

const fetchUsers = async () => {
  const filters: FilterArgs = {}
  if (payload.value.user !== '') {
    const spl = payload.value.user.split(' ')
    
  }
  await identities.fetchUsers({
    page: 1,
    itemsPerPage: 20,
    sortBy: "firstname",
  }, filters)
  identities.users.forEach((item) => {
    existingUsers.value.push({ value: item.id, title: `${item.firstname} ${item.lastname}` })
  })
}


const existingUsers = ref<{ value: string, title:string} []>([])
onMounted(async () => {
  if (dialogStore.item && Object.keys(dialogStore.item).length > 0) {
    payload.value = dialogStore.item
  } else {
    dialogStore.item = payload.value
  }

  if (!props.sameUser) {
    await identities.fetchUsers()
    identities.users.forEach((item) => {
      existingUsers.value.push({ value: item.id, title: `${item.firstname} ${item.lastname}` })
    })
  } else {
    payload.value.user = authStore.userID
  }


  dialogStore.registerCallback(async () => {
    const keyPair = generateKeyPair()
    const validation = [] as (string | boolean)[]
    validation.push(validationRules.required(payload.value.name, 'name'))

    const errors = [] as string[]
    validation.forEach((v) => {
      if (typeof v === 'string') {
        errors.push(v)
      }
    })
    if (errors.length > 0) {
      throw new ValidationError(errors)
    }

    const data = await useDataStore().create({
      ...payload.value,
      publicKey: b64encode(keyPair.publicKey),
    })

    dialogStore.hideOk()
    peerConfig.value = {
      ...data,
      privateKey: b64encode(keyPair.secretKey),
    } as PeerConfig
    
    console.log(peerConfig.value)
    created.value = true

    return true
  })
})


const payload = ref({
  name: '',
  type: 'laptop',
  dns: [] as string[],
  user: '',
  description: '',
  endpoint: '',
  keepAlive: true,
  allowedIps: [],
  publicKey: '',
} as CreateDeviceRequest)

const peerConfig = ref({})
type PeerConfig = {
  id: number;
  name: string;
  type: string;
  dns: string[];
  user?: string | undefined;
  description?: string | undefined;
  publicKey: string;
  presharedKey: string;
  endpoint: string;
  allowedIps: string;
  privateKey: string;
  keepAlive: boolean;
}

const getPeerConfig = (peerConfig: any) => {
  return `[Interface]
PrivateKey = ${peerConfig.privateKey}
Address = ${peerConfig.endpoint}
DNS = ${peerConfig.dns.join(",")}

[Peer]
PublicKey = ${peerConfig.publicKey}
PresharedKey = ${peerConfig.presharedKey}
AllowedIPs = ${peerConfig.allowedIps}
Endpoint = ${(new URL(API_URL)).hostname}:51820
PersistentKeepalive = ${peerConfig.keepalive}`
}

</script>
