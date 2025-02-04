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
import user from "./system/user";
import department from "./system/department";
import message from "./system/messages";

import category from "./shop/category"
import brand from "./shop/brand";
import products from "./shop/products";
import order from "./shop/order";
import attribute from "./shop/attribute";


import role from "./system/role";

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
        name="message"
        {...message}
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


      <Resource
        name="shopcategory"
        {...category}
      />
      <Resource
        name="attribute"
        {...attribute}
      />
      <Resource
        name="brand"
        {...brand}
      />
      <Resource
        name="products"
        {...products}
      />
      <Resource
        name="order"
        {...order}
      />

    </Admin>


);
