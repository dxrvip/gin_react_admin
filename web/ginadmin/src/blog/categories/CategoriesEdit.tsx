import {Edit, SimpleForm, useRecordContext, TextInput} from "react-admin"


const CategoriesEdit = () => {
    const CategoryTitle = () => {
        const record = useRecordContext();
        return <span>编辑: {record ? record.name : ''}</span>;
    }
    return (
        <Edit title={<CategoryTitle />} >

            <SimpleForm>
                
                    <TextInput source="id" label="ID"  disabled/>
                    <TextInput source="name" label="名称" />
                
            </SimpleForm>
        </Edit>
    )
}



export default CategoriesEdit