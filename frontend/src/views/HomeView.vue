<script setup lang="ts">
import { useI18n } from "vue-i18n";

import Player from "@/components/Player.vue";
import { onMounted, reactive } from "vue";

import { StartStopGame, EndRound } from "../../wailsjs/go/main/GameData.js" 
import { EventsOn } from "../../wailsjs/runtime/runtime.js"

const { t } = useI18n();

const data = reactive({
  player1: "",
  player2: "",
  game_data: "",
  game_status: t("game.start"),
})

EventsOn("game_data", (res) => {
  data.game_data = res
  data.player1 = res["Players"][0]
  data.player2 = res["Players"][1]
})

function startGame() {
  StartStopGame()
  if (data.game_status == t("game.start") || data.game_status == t("game.startNew")) {
    data.game_status = t("game.restart")
  } else if (data.game_status == t("game.restart")){
    data.game_status = t("game.startNew")
  }
}

function endRound() {
  EndRound()
}

onMounted(() => {
  
})
</script>

<template>
  <div class="home">
    <!-- Logo -->
    <!-- <p>{{ data.game_data }}</p> -->
    <div class="grid grid-cols-2 mt-5 h-100" v-if="data.player1 && data.player2">
      <Player :player="data.player1" @click="endRound" />
      <Player :player="data.player2" @click="endRound" />
    </div>
    <div class="w-100 my-10 flex justify-center">
      <button class="button mx-auto text-xl text-white" @click="startGame">{{ data.game_status }}</button>
    </div>
  </div>
</template>

<style lang="scss">
.button {
  background-color: rgba(0, 0, 0, 0.147);
  padding: 10px;
  border-radius: 10px;
}
.home {
  .link {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: center;
    margin: 18px auto;
    .btn {
      display: block;
      width: 150px;
      height: 50px;
      line-height: 50px;
      padding: 0 5px;
      margin: 0 30px;
      border-radius: 8px;
      text-align: center;
      font-weight: 700;
      font-size: 16px;
      white-space: nowrap;
      text-decoration: none;
      cursor: pointer;
      &.start {
        background-color: #fd0404;
        color: #ffffff;
        &:hover {
          background-color: #ec2e2e;
        }
      }
      &.star {
        background-color: #ffffff;
        color: #fd0404;
        &:hover {
          background-color: #f3f3f3;
        }
      }
    }
  }
}
</style>
