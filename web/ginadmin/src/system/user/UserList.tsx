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
                <TextField source="title" />
                <DateField source="published_at" />
                <TextField source="category" />
                <BooleanField source="commentable" />
            </Datagrid>
        </List>
    );
}
