import {
    Create,
    DateInput,
    email,
    required,
    SimpleForm,
    TextInput,
    maxValue,
    SelectInput,
} from 'react-admin'
const choices = [
    { id: 'other', name: '其他' },
    { id: 'female', name: '女' },
    { id: 'male', name: '男' },
 ];
function UserCreate() {
    return (
        <Create>
            <SimpleForm>
                <TextInput source="userName" validate={[required()]} label="账号" />
                <TextInput type='password' source="password" label="密码" defaultValue={new Date()} />
                <TextInput source="nikeName" validate={[required(), maxValue(50)]} label="昵称" />
                <TextInput source="email" validate={[required(), email()]} label="邮箱" />
                <SelectInput source="gender" choices={choices} label="性别" defaultValue={"other"} />
            </SimpleForm>
        </Create>
    );
}

export default UserCreate;