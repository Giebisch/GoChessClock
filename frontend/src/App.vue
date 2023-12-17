<script setup lang="ts">
import { useI18n } from "vue-i18n";

const { t, availableLocales: languages, locale } = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};

const onclickMinimise = () => {};

const onclickQuit = () => {};

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <!-- Header -->
  <div class="header mt-7">
    <!-- navigation -->
    <div class="nav">
      <router-link to="/">{{ t("nav.home") }}</router-link>
      <router-link to="/statistics">{{ t("nav.statistics") }}</router-link>
    </div>
    <!-- Menu -->
    <div class="menu">
      <div class="language">
        <div
          v-for="item in languages"
          :key="item"
          :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)"
          class="lang-item"
        >
          {{ t("languages." + item) }}
        </div>
      </div>
    </div>
  </div>
  <!-- Page -->
  <div class="view">
    <router-view />
  </div>
</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  background-color: rgba(160, 160, 160, 0.367);
  overflow: hidden;
}
.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: rgba(0, 0, 0, 0.109);
  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #0000001c;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;
      &:hover,
      &.router-link-exact-active {
        background-color: #00000038;
        color: #ffffff;
      }
    }
  }
  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;
    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #c3c3c333;
      overflow: hidden;
      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;
        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }
        &.active {
          background-color: #ff050542;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }
    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;
      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #ac7edc4b;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;
        &:hover {
          background-color: #d7a8d8bc;
          color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}

</style>
