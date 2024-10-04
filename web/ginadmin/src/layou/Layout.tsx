import type { ReactNode } from "react";
import { Layout as RALayout, 
  CheckForApplicationUpdate, 
  AppBar, 
  Logout, 
  UserMenu, 
  TitlePortal,
  MenuItemLink } from "react-admin";
import Menu from "./Menu";
import SettingsIcon from "@mui/icons-material/Settings"

const MyUserMenu = (props: any) => {

  return (
    <UserMenu {...props}>
      <MenuItemLink
        to="/profile"
        primaryText="设置"
        leftIcon={<SettingsIcon />}
      />
      <Logout key="logout" />
    </UserMenu>
  )
}

export const MyAppBar = () => (
  <AppBar userMenu={<MyUserMenu />}>
      <TitlePortal variant="h6" component="h5" />
      {/* <SettingsButton /> */}
  </AppBar>
);


export const Layout = ({ children }: { children: ReactNode }) => (
  <RALayout menu={Menu} appBar={MyAppBar} >
    {children}
    <CheckForApplicationUpdate />
  </RALayout>
);
