import { 
    List, 
    SearchInput,
    TextInput,
    Datagrid, 
    TextField, 
    DateField, 
    ReferenceField,
    DateInput,
} from 'react-admin';

const postFilters = [
    <SearchInput source="q" alwaysOn />,
    <TextInput label="Title" source="title" defaultValue="Hello, World!" />,
    <TextInput label="Desc" source="desc" defaultValue="Hello, World!" />,
    <TextInput label="分类" source="cid" defaultValue="Hello, World!" />,
    <DateInput label="添加日期" source="CreatedAt" defaultValue={Date.now()} />
]


const ArticleList = () => (
    <List filters={postFilters}>
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

export default ArticleList;