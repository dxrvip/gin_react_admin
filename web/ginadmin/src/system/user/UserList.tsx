import {
    List,
    Datagrid,
    TextField,
    DateField,
    BooleanField,
} from "react-admin"


export function UserList(prop: any) {
    return (
        <List>
            <Datagrid>
                <TextField source="id" />
                <TextField source="username" />
                <DateField source="nike_name" />
                <TextField source="active" />
                <BooleanField source="gender" />
            </Datagrid>
        </List>
    );
}
