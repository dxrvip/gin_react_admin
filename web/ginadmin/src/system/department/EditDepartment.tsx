import {
    Edit,
    SimpleForm,
    TextInput,
    ReferenceInput,
    SelectInput,
    useRecordContext,
} from "react-admin"
const ParentIdInput = (props: any) => {
    const record = useRecordContext()
    if (!record) return null
    return (
        <>
            <ReferenceInput source="parent_id" reference="department" label="父部门" filter={{ "id_ne": record.id }}>
                <SelectInput optionText="name" />
            </ReferenceInput>
        </>
    )
}
function EditDepartment() {
    return (
        <Edit title="Edit Department" >
            <SimpleForm>
                <TextInput source="name" label="部门名称" />
                {/* <ReferenceInput source="parent_id" reference="department">
                    <SelectInput optionText="name" label="父部门" />
                </ReferenceInput> */}
                <ParentIdInput />
            </SimpleForm>
        </Edit>
    );
}

export default EditDepartment;