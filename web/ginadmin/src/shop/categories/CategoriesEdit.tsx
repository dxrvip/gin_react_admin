import {
    Edit,
    TextInput,
    SimpleForm,
    ReferenceInput,
    SelectInput,
    required,
    minLength,
    maxLength,
    useRecordContext,
} from "react-admin"


function ParentInput(props: any) {
    const record = useRecordContext();  // 获取当前记录
    const currentId = record?.id;      // 获取当前 ID
    console.log(currentId)
    return (
        <>
            <ReferenceInput source="parentId" reference="categories" filter={{ id_ne: currentId }} >
                <SelectInput label="父分类" resettable />
            </ReferenceInput>

        </>
    )
}

function CategoriesEdit(props: any) {


    return (
        <Edit title="Edit Product Categories" {...props}>
            <SimpleForm>
                <TextInput source="name" label="分类名称" validate={[required(), minLength(2), maxLength(100)]} />
                <TextInput source="description" label="分类描述" validate={[maxLength(500)]} />
                <ParentInput />
            </SimpleForm>
        </Edit>
    );
}

export default CategoriesEdit;