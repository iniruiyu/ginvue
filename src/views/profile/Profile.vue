<template>
  <h3>
    您好 <span v-if="userInfo">{{ userInfo.name }}</span>
  </h3>
</template>

<script setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
const store = useStore()
const userInfo = computed(() => {
  return store.state.userModule.userInfo
})
</script>
