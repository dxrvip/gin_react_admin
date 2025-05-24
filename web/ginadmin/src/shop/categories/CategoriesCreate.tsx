import {Create, TextInput,SimpleForm, ReferenceInput, SelectInput, required, minLength, maxLength} from "react-admin"

function CategoriesCreate(props: any) {
    return (
        <Create title="创建产品类别" {...props} redirect="list">
        <SimpleForm>
            <TextInput source="name" label="分类名称" validate={[required(), minLength(2), maxLength(100)]} />
            <TextInput source="description" label="分类描述" validate={[maxLength(500)]} />
            <ReferenceInput source="parentId" reference="categories" >
                <SelectInput label="父分类" resettable/>
            </ReferenceInput>
        </SimpleForm>
    </Create>
    );
}

export default CategoriesCreate;