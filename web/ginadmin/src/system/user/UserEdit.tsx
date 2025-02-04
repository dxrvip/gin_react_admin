import {
  createContext,
  useState,
  useCallback,
  useMemo,
  useContext,
} from "react";
import {
  TextInput,
  SimpleForm,
  required,
  useNotify,
  SelectInput,
  SaveContextProvider,
  useGetIdentity,
  useRedirect,
  Toolbar,
  SaveButton,
  maxLength,
  minLength,
  BooleanInput,
  email,
  useUpdate,
  Edit,
} from "react-admin";
import { choices } from "./UserCreate";
const ProfileContext = createContext({
  profileVersion: 0,
  refreshProfile: () => { },
});

export const ProfileProvider = ({ children }: { children: any }) => {
  const [profileVersion, setProfileVersion] = useState(0);
  const context = useMemo(
    () => ({
      profileVersion,
      refreshProfile: () => {
        setProfileVersion((currentVersion) => currentVersion + 1);
      },
    }),
    [profileVersion]
  );

  return (
    <ProfileContext.Provider value={context}>
      {children}
    </ProfileContext.Provider>
  );
};

export const useProfile = () => useContext(ProfileContext);

const CustomToolbar = (props: any) => (
  <Toolbar {...props}>
    <SaveButton />
  </Toolbar>
);

export const UserEdit = ({ ...props }) => {
  // const notify = useNotify();
  const redirect = useRedirect();
  const [saving, _] = useState(false);
  const { isLoading: isUserIdentityLoading, data } = useGetIdentity();
  if (!isUserIdentityLoading && !data?.username) {
    redirect("/login");
  }


  if (isUserIdentityLoading) {
    return null;
  }

  return (
    <Edit>
      <SimpleForm
        sx={{ width: 699 }}
        toolbar={<CustomToolbar />}
      >
        <TextInput source="id" label="用户ID" disabled />
        <TextInput
          source="username"
          label="登陆名称"
          validate={[required(), maxLength(10), minLength(6)]}
        />
        <TextInput source="nike_name" label="昵称" validate={required()} />
        <TextInput source="email" label="邮箱" validate={[required(), email()]} />
        <SelectInput
          source="gender"
          choices={choices}
          label="性别"
          defaultValue={"未知"}
        />
        <BooleanInput source="status" label="状态" />
      </SimpleForm>
    </Edit>
  );
};
