import { createI18n } from "vue-i18n";

import de from "./locales/de.json";
import en from "./locales/en.json";

const i18n = createI18n({
  locale: "de",
  fallbackLocale: "de",
  legacy: false,
  messages: {
    de: de,
    en: en,
  },
});

export default i18n;
