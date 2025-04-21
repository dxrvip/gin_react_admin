import {
    Edit
} from "react-admin";
import AttributeForm from "./AttributeForm";



function AttributeEdit(props: any) {
    return (
        <Edit title="Edit Attribute Value" {...props}>
            <AttributeForm />
        </Edit>
    );
}

export default AttributeEdit;