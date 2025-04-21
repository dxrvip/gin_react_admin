
import {List, Datagrid, TextField, DateField} from "react-admin"

function CategoriesList(props: any) {
    return (
        <List>

            <Datagrid>
                <TextField source="id" />
                <TextField source="name" label="名称" />
                <DateField source="createAd" label="创建时间" />
                <TextField source="description" label="描述"/>
            </Datagrid>
        </List>
    );
}

export default CategoriesList;