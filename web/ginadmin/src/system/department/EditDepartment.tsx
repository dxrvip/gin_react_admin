import {
    Edit,
    SimpleForm,
    TextInput,
    ReferenceInput,
    SelectInput,
} from "react-admin"

function EditDepartment() {
    return (
        <Edit title="Edit Department" >
            <SimpleForm>
                <TextInput source="name" label="部门名称" />
                <ReferenceInput source="parent_id" reference="department">
                    <SelectInput optionText="name" label="父部门" />
                </ReferenceInput>
            </SimpleForm>
        </Edit>
    );
}

export default EditDepartment;