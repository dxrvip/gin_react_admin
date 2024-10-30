import { useEffect, useState } from "react"
import {
    List,
    Datagrid,
    TextField,
    NumberField,
    EditButton,
    DeleteButton,
    BooleanField,
    useGetList,
    Loading,
    useRecordContext,
    useUpdate,
} from 'react-admin'
import { Button } from "@mui/material"
import RouteIcon from '@mui/icons-material/Route';
import DialogWindow from "../../components/Dialogwindow"
import IndeterminateCheckbox from "../../components/IndeterminateCheckbox"
import { PackageItem } from "../../utils/dataProvider";

const DialogContent = (props: { setMenuDatas: (d: any) => void }) => {
    const { setMenuDatas } = props;

    const { data, isPending, error } = useGetList(
        'systemMenu'
    );

    if (isPending) { return <Loading />; }
    if (error) { return <p>菜单数据获取失败</p>; }

    return <IndeterminateCheckbox data={data as any} setMenuDatas={setMenuDatas} />


}
const DialogActions = (props: { menuDatas: PackageItem[], setOpen: (b: boolean) => void }) => {
    const record = useRecordContext()
    const [menus, setMenus] = useState<string[]>([])
    const { menuDatas, setOpen } = props
    const [update, { isPending, error }] = useUpdate(
        'role',
        { id: record?.id, data: { menus: JSON.stringify(menus) }, previousData: record }
    )
    useEffect(() => {
        if (menus.length > 0) {
            console.log('准备提交数据:', menus);
            update()
        }
    }, [menus, update])
    if (!record) return null;
    // 提交更新

    const handleClick = () => {

        const newMenu = menuDatas.flatMap(element =>
            element.func.filter(node => node.active).map(node => node.name)
        );
        setMenus(newMenu)

        console.log(menuDatas, record, newMenu, menus)
    }
    return (
        <>
            <Button onClick={handleClick}>保存权限</Button>
            <Button autoFocus onClick={() => setOpen(false)}>
                取消
            </Button>
        </>
    )
}
const ButtonGroupFiled = (props: any) => {
    const [open, setOpen] = useState<boolean>(false)
    const [menuDatas, setMenuDatas] = useState<PackageItem[]>()

    return (
        <>
            <EditButton label='编辑' />
            <DeleteButton label='删除' />
            <Button variant="text" endIcon={<RouteIcon />} onClick={() => {
                setOpen(true)
            }}>
                权限编辑
            </Button>
            <DialogWindow onClose={setOpen} open={open} dialogActions={<DialogActions setOpen={setOpen} menuDatas={menuDatas as PackageItem[]} />}>
                <DialogContent setMenuDatas={setMenuDatas} />
            </DialogWindow>
        </>
    )
}



function ListRole() {
    return (
        <List title="角色列表">
            <Datagrid rowClick={false}>
                <TextField source='id' label="ID" />
                <TextField source="name" label="名称" />
                <TextField source="key" label="权限标识符" />
                <NumberField source="sort" label="权限排序" />
                <BooleanField source="active" label="状态" />
                <ButtonGroupFiled label="操作" />
            </Datagrid>
        </List>
    );
}

export default ListRole;