import { List, Datagrid, TextField, DateField, ReferenceField } from 'react-admin';

const PostList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <TextField source="title" />
            <DateField source="CreatedAt" label="添加日期" />
            <TextField source="desc" label="摘要"/>
            <ReferenceField source="cid" reference="category" label="分类" >
                <TextField source="name" />
            </ReferenceField>
        </Datagrid>
    </List>
);

export default PostList;