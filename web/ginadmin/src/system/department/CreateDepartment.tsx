import {
    Create,
    SimpleForm,
    TextInput,
    ReferenceInput,
    SelectInput
} from 'react-admin'

function CreateDepartement() {
    return (
        <Create>
            <SimpleForm>
                <TextInput source="name" label="部门名称" />
                <ReferenceInput source="parent_id" reference="department" label="父部门">
                    <SelectInput optionText="name" />
                </ReferenceInput>
            </SimpleForm>
        </Create>
    );
}

export default CreateDepartement;