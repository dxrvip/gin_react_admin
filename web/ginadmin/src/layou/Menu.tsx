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
import system from "../system"
import post from "../blog/article"
import department from "../system/department";
import user from "../system/user"
import role from "../system/role"
import news from "../system/news"
import categories from "../blog/categories"
type MenuName = 'blog' | 'system';


const Menu = ({ dense = false }: MenuProps) => {
    const [state, setState] = useState({
        blog: false,
        system: true,
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
                handleToggle={() => handleToggle('system')}
                isOpen={state.system}
                name="系统管理"
                icon={<system.icon />}
                dense={dense}
            >
                <MenuItemLink
                    to="/department"
                    state={{ _scrollToTop: true }}
                    primaryText="部门管理"
                    leftIcon={<department.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/role"
                    state={{ _scrollToTop: true }}
                    primaryText="角色管理"
                    leftIcon={<role.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/user"
                    state={{ _scrollToTop: true }}
                    primaryText="用户管理"
                    leftIcon={<user.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/news"
                    state={{ _scrollToTop: true }}
                    primaryText="消息中心"
                    leftIcon={<news.icon />}
                    dense={dense}
                />
            </SubMenu>
            <SubMenu
                handleToggle={() => handleToggle('blog')}
                isOpen={state.blog}
                name="博客管理"
                icon={<blog.blogIcon />}
                dense={dense}
            >
                <MenuItemLink
                    to="/article"
                    state={{ _scrollToTop: true }}
                    primaryText="文章管理"
                    leftIcon={<post.icon />}
                    dense={dense}
                />
                <MenuItemLink
                    to="/category"
                    state={{ _scrollToTop: true }}
                    primaryText="分类管理"
                    leftIcon={<categories.icon />}
                    dense={dense}
                />
            </SubMenu>
           
        </Box>
    )
}

export default Menu;