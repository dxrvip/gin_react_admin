// import { BooleanField, Datagrid, DeleteButton, EditButton, List, ShowButton, TextField } from "react-admin";

// function AttributeList(props: any) {
//     return ( 
//         <List {...props}>
//         <Datagrid>
//             <TextField source="id" label="ID" />
//             <TextField source="name" label="属性名称" />
//             <TextField source="category_name" label="绑定的分类" />
//             <TextField source="value_name" label="预设值" />
//             <BooleanField source="is_required" label="是否必填" />

//         </Datagrid>
//     </List>
//      );
// }


// src/views/attributes/AttributeList.tsx
import { List, Datagrid, TextField, BooleanField, EditButton, DeleteButton, ArrayField, SingleFieldList, ChipField } from 'react-admin';

const AttributeList = () => (
    <List>
        <Datagrid rowClick="edit">
            <TextField source="id" />
            <TextField source="name" label="属性名称" />
            <TextField source="type" label="类型" />
            <BooleanField source="isRequired" label="必填" />

            <ArrayField source="categories" label="关联分类">
                <SingleFieldList>
                    <ChipField source="name" size="small" />
                </SingleFieldList>
            </ArrayField>

            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
);
export default AttributeList;