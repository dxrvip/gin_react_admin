import {
  Admin,
  Resource,
  CustomRoutes,
} from "react-admin";
import { Route } from "react-router";
import { Layout, Component } from "./layou";
import dataProvider from "./utils/dataProvider";
import { authProvider } from "./utils/authProvider";

import i18nProvider from "./utils/i18nProvider"
import article from "./blog/article"
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
      name="article"
      {...article}
    />

    <Resource
      name="category"
      {...categories}
    />

  </Admin>
);
