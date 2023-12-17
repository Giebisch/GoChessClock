<script setup lang="ts">
defineProps<{
  player: any;
}>();
</script>

<template>
  <div class="w-100 text-center py-20 background text-white" :class="{'turn': player.Turn}">
    <p class="text-xl">{{ player.Name }}</p>
    <p class="font-bold text-8xl my-20">{{ player.RemainingTime }}</p>
  </div>
</template>

<style scoped>

.background {
  border-radius: 10px;
  margin: 15px;
}
.turn {
  background-color: rgba(0, 0, 0, 0.147);
}
</style>
