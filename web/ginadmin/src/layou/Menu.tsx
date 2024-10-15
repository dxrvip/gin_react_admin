import { useState } from "react";
import {
    MenuProps,
    useSidebarState,
    DashboardMenuItem,
    MenuItemLink,
} from "react-admin";
import { Box } from "@mui/material";
import SubMenu from "./SubMenu";

import * as blog from "../blog"
import post from "../blog/article"
import categories from "../blog/categories"
type MenuName = 'menuCatalog' | 'menuSales' | 'menuCustomers';


const Menu = ({ dense = false }: MenuProps) => {
    const [state, setState] = useState({
        menuCatalog: true,
        menuSales: true,
        menuCustomers: true,
    });

    const [open] = useSidebarState();
    const handleToggle = (menu: MenuName) => {
        setState(state => ({ ...state, [menu]: !state[menu] }));
    };

    return (
        <Box
            sx={{
                width: open ? 200 : 50,
                marginTop: 1,
                marginBottom: 1,
                transition: theme =>
                    theme.transitions.create('width', {
                        easing: theme.transitions.easing.sharp,
                        duration: theme.transitions.duration.leavingScreen,
                    }),
            }}
        >
            <DashboardMenuItem />
            <SubMenu
                handleToggle={() => handleToggle('menuSales')}
                isOpen={state.menuSales}
                name="系统管理"
                icon={<blog.blogIcon />}
                dense={dense}
            >
                <MenuItemLink
                    to="/article"
                    state={{ _scrollToTop: true }}
                    primaryText="部门管理"
                    leftIcon={<post.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="角色管理"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="用户管理"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="消息中心"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
            </SubMenu>
            <SubMenu
                handleToggle={() => handleToggle('menuSales')}
                isOpen={state.menuSales}
                name="博客管理"
                icon={<blog.blogIcon />}
                dense={dense}
            >
                <MenuItemLink
                    to="/article"
                    state={{ _scrollToTop: true }}
                    primaryText="文章"
                    leftIcon={<post.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="分类"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
            </SubMenu>
            <SubMenu
                handleToggle={() => handleToggle('menuSales')}
                isOpen={state.menuSales}
                name="商城管理"
                icon={<blog.blogIcon />}
                dense={dense}
            >
                <MenuItemLink
                    to="/article"
                    state={{ _scrollToTop: true }}
                    primaryText="商品管理"
                    leftIcon={<post.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="订单管理"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="商品分类"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />

                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="属性分类"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />

                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="物流管理"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />


            </SubMenu>
        </Box>
    )
}

export default Menu;