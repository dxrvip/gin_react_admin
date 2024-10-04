import {
  Admin,
  Resource,
  CustomRoutes,
} from "react-admin";
import { Route } from "react-router";
import { Layout, Component } from "./layou";
import dataProvider from "./dataProvider";
import { authProvider } from "./authProvider";

import i18nProvider from "./i18nProvider"
import post from "./blog/post"
import categories from "./blog/categories"
import { ProfileEdit } from "./user";
export const App = () => (
  <Admin
    layout={Layout}
    dashboard={Component}
    i18nProvider={i18nProvider}
    dataProvider={dataProvider}
    authProvider={authProvider}
  >
    <CustomRoutes>
      <Route path="/profile" element={<ProfileEdit />} />
    </CustomRoutes>
  
    <Resource
      name="posts"
      {...post}
    />

    <Resource
      name="category"
      {...categories}
    />

  </Admin>
);
