import {
  Admin,
  Resource,
  ListGuesser,
  EditGuesser,
  ShowGuesser,
} from "react-admin";
import { Layout } from "./layou";
import dataProvider from "./dataProvider";
import { authProvider } from "./authProvider";

import i18nProvider from "./i18nProvider"
import post from "./blog/post"
import categories from "./blog/categories"
export const App = () => (
  <Admin
    layout={Layout}
    i18nProvider={i18nProvider}
    dataProvider={dataProvider}
    authProvider={authProvider}
  >
    <Resource
      name="posts"
      {...post}
    />

    <Resource
      name="category"
      {...categories}
    />
    <Resource
      name="users"
      options={{ label: "用户管理" }}
      list={ListGuesser}
      edit={EditGuesser}
      show={ShowGuesser}
    />
  </Admin>
);
