import { Create, SimpleForm, TextInput, DateInput, required } from 'react-admin';

const CategoriesCreate = () => (
    <Create title="分类">
        <SimpleForm>
            <TextInput source="name" label="名称" validate={[required()]} />
        </SimpleForm>
    </Create>
);


export default CategoriesCreate;