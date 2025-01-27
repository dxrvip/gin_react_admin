import {
  Create,
  email,
  required,
  SimpleForm,
  TextInput,
  maxValue,
  SelectInput,
  maxLength,
} from "react-admin";
export const choices = [
  { id: "other", name: "其他" },
  { id: "female", name: "女" },
  { id: "male", name: "男" },
];

//验证密码是否正确

const passwordValidation = (value: any, allValues: any) => {
  if (value === allValues.password) {
    return undefined;
  }
  return "二次密码不一致！";
};

function UserCreate() {
  return (
    <Create redirect="list">
      <SimpleForm>
        <TextInput source="userName" validate={[required()]} label="账号" />
        <TextInput
          type="password"
          source="password"
          label="密码"
          validate={[required(), maxLength(20)]}
        />
        <TextInput
          type="password"
          source="re_password"
          label="确认密码"
          validate={[required(), maxLength(20), passwordValidation]}
        />
        <TextInput
          source="nikeName"
          validate={[required(), maxValue(50)]}
          label="昵称"
        />
        <TextInput
          source="email"
          validate={[required(), email()]}
          label="邮箱"
        />
        <SelectInput
          source="gender"
          choices={choices}
          label="性别"
          defaultValue={"other"}
        />
      </SimpleForm>
    </Create>
  );
}

export default UserCreate;
