import {
    Edit
} from "react-admin";
import AttributeForm from "./AttributeForm";



function AttributeEdit(props: any) {
    return (
        <Edit title="编辑属性" {...props} redirect="list">
            <AttributeForm />
        </Edit>
    );
}

export default AttributeEdit;