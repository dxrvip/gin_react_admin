import {
    List,
    Datagrid,
    TextField,
    ChipField,
    BooleanField,
    EditButton,
    DeleteButton,
} from "react-admin"

const GroupBut = (props: any) => {
    return (
        <>
            <EditButton />
            <DeleteButton />
        </>
    )
}

//male', 'female', 'other

export function UserList(prop: any) {
    return (
        <List>
            <Datagrid>
                <TextField source="id" label="ID"/>
                <TextField source="username" label="用户名" />
                <TextField source="nike_name" label="昵称" />
                <ChipField source="gender" label="性别" />
                <BooleanField source="active" label="状态" />
                <GroupBut label="操作" />
            </Datagrid>
        </List>
    );
}
