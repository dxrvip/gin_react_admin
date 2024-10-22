


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
    SaveContextProvider,
    useGetIdentity,
    useRedirect,
    Toolbar,
    SaveButton,
    maxLength,
    minLength,
} from "react-admin";
//   import { userApi } from "../providers/env";

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

export const ProfileEdit = ({ ...props }) => {
    const notify = useNotify();
    const redirect = useRedirect();
    const [saving, setSaving] = useState(false);
    const { refreshProfile, profileVersion } = useProfile();
    const { isLoading: isUserIdentityLoading, data } = useGetIdentity();
    console.log("data", data,refreshProfile, profileVersion)
    if (!isUserIdentityLoading && !data?.username) {
        redirect("/login");
    }

    const handleSave = useCallback(
        (values: any) => {
            console.log("values", values)
        },
        [notify, refreshProfile, redirect]
    );

    if (isUserIdentityLoading) {
        return null;
    }

    return (
        <SaveContextProvider
            value={{ save: handleSave, saving }}
            key={profileVersion}
        >
            <SimpleForm sx={{ width: 699}} record={data ? data : {}} toolbar={<CustomToolbar />}>
                <TextInput source="user_id" label="用户ID" disabled />
                <TextInput source="username" label="登陆名称" validate={[required(), maxLength(10), minLength(6)]} />
                <TextInput source="full_name" label="昵称" validate={required()} />
            </SimpleForm>
        </SaveContextProvider>
    );
};