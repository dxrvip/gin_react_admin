import polyglotI18nProvider from 'ra-i18n-polyglot';
import chineseMessages from '@haxqer/ra-language-chinese';
import en from 'ra-language-english';



const i18nProvider = polyglotI18nProvider(locale =>
    locale === 'zh_CN' ? chineseMessages : en,
    'zh_CN', // Default locale
    [
        { locale: 'en', name: 'English' },
        { locale: 'zh_CN', name: '中文简体' }
    ],
    { allowMissing: true }
);




export default i18nProvider;