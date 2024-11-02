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
    useNotify,
    SimpleForm,
    Edit,
    SelectArrayInput,
    SaveButton,
    Toolbar,
    useRedirect,
} from 'react-admin'
import LoadingButton from '@mui/lab/LoadingButton';
import { Box, Button } from "@mui/material"
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
    const notify = useNotify()
    const { menuDatas, setOpen } = props
    const [update, { isPending, error }] = useUpdate(
        'role',
        { id: record?.id, data: { menus: menus }, previousData: record },
        {
            onSuccess: (val => {
                notify('权限编辑成功！', { type: "success" })
            }),
            onError: (val => {
                notify(val?.message, { type: "error" })
            })
        }

    )
    useEffect(() => {
        if (menus.length > 0) {
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

    }
    return (
        <>
            <LoadingButton loading={isPending} onClick={handleClick}>保存权限</LoadingButton>
            <Button autoFocus onClick={() => setOpen(false)}>
                关闭
            </Button>
        </>
    )
}
const PostSaveButton = () => {
    const notify = useNotify();
    const redirect = useRedirect();
    const onSuccess = (data: any) => {
        notify(`Post "${data.title}" saved!`);
        redirect('/posts');
    };
    return (
        <SaveButton label="提交" type="button" mutationOptions={{ onSuccess }} />
    );
};
export const MyToolbar = () => (
    <Toolbar>
        <PostSaveButton />
    </Toolbar>
);
const DialogAddUserContent = (props: any) => {
    const { data, isPending, error } = useGetList(
        'user'
    );
    if (isPending) return <Loading />

    if (error) return null
    console.log(data)
    return (
        <Box sx={{width: 300 }}>

            <Edit id={2} redirect="list">
                <SimpleForm sx={{p:2, m: 3}} toolbar={<MyToolbar />}>
                    <SelectArrayInput source="userId" choices={data} optionText="username" />

                </SimpleForm>
            </Edit>
        </Box>


    )
}

const ButtonGroupFiled = (props: any) => {
    const [open, setOpen] = useState<boolean>(false)
    const [addUserOpen, setAddUserOpen] = useState<boolean>(false)
    const [menuDatas, setMenuDatas] = useState<PackageItem[]>()

    return (
        <>
            <EditButton label='编辑' />
            <Button variant="text" endIcon={<RouteIcon />} onClick={() => setOpen(true)}>
                权限编辑
            </Button>
            <Button variant="text" onClick={() => setAddUserOpen(true)}>添加用户</Button>
            <DeleteButton label='删除' />
            <DialogWindow onClose={setOpen} open={open} dialogActions={<DialogActions setOpen={setOpen} menuDatas={menuDatas as PackageItem[]} />}>
                <DialogContent setMenuDatas={setMenuDatas} />
            </DialogWindow>
            {/* 添加用户窗口 */}
            <DialogWindow open={addUserOpen} onClose={setAddUserOpen}>
                <DialogAddUserContent />
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