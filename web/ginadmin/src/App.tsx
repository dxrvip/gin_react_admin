import {
  Admin,
  Resource,
  CustomRoutes,
} from "react-admin";
import { ThemeProvider, CssBaseline } from '@mui/material';
import { Route } from "react-router";
import { Layout, Component } from "./layou";
import dataProvider from "./utils/dataProvider";
import { authProvider } from "./utils/authProvider";

import i18nProvider from "./utils/i18nProvider"
import article from "./blog/article"
import categories from "./blog/categories"
import user from "./system/user";
import department from "./system/department";
import news from "./system/news";
import role from "./system/role";
import theme from "./utils/theme";

export const App = () => (

    <Admin
      layout={Layout}
      dashboard={Component}
      i18nProvider={i18nProvider}
      dataProvider={dataProvider}
      authProvider={authProvider}
    >
      <CustomRoutes>
        <Route path="/profile" element={<user.edit />} />
      </CustomRoutes>
      <Resource
        name="user"
        {...user}
      />
      <Resource
        name="role"
        {...role}
      />
      <Resource
        name="news"
        {...news}
      />
      <Resource
        name="department"
        {...department}
      />
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
