import { Datagrid, List, TextField } from "react-admin";

function BrandList(props: any) {
    return (
        <List {...props}>
            <Datagrid>
                <TextField source="id" label="ID" />
                <TextField source="name" label="品牌名称" />
                <TextField source="logo" label="Logo URL" />
                <TextField source="description" label="品牌描述" />

            </Datagrid>
        </List>
    );
}

export default BrandList;