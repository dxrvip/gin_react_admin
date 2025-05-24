import { Create, maxLength, minLength, required, SimpleForm, TextInput } from "react-admin";

function BrandCreate(props: any) {
    return (
        <Create title="创建一个品牌" {...props} redirect="list">
            <SimpleForm>
                <TextInput source="name" label="品牌名称" validate={[required(), minLength(2), maxLength(255)]} />
                <TextInput source="logo" label="Logo URL" />
                <TextInput source="description" label="品牌描述" />
            </SimpleForm>
        </Create>
    );
}

export default BrandCreate;