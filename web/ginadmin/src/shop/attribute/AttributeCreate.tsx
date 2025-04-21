// import {
//     Create,
//     ReferenceInput,
//     required,
//     SelectInput,
//     SelectArrayInput,
//     SimpleForm,
//     TextInput,
//     useCreateSuggestionContext,
//     useCreate,
//     useNotify,

// } from "react-admin";
// import { useFormContext } from "react-hook-form"
// import {
//     TextField as MuiTextField,
//     Button,
//     Dialog,
//     DialogActions,
//     DialogContent,
// } from '@mui/material';
// import React from "react";


// function AttributeCreate(props: any) {

//     return (
//         <Create title="Create an Attribute Value" {...props}>
//             <SimpleForm>
//                 <ReferenceInput source="attrbutekey_id" reference="attrbutekey">
//                     <SelectInput
//                         label="属性键"
//                         validate={required()}
//                         create={<CreateDialog children={<CreateCategoriesKeyFrom />} />}
//                     />
//                 </ReferenceInput>


//                 <ValueInput />
//                 <SelectInput source="is_required" choices={[{ id: true, name: "必填" }, { id: false, name: "不强求" }]} label="是否必填" validate={[required()]} />
//                 <ReferenceInput source="category_id" reference="categories">

//                     <SelectArrayInput label="绑定分类" />
//                 </ReferenceInput>
//             </SimpleForm>
//         </Create>
//     );
// }

// export default AttributeCreate;


// export const ValueInput = () => {

//     const form = useFormContext()

//     return (
//         <>

//             <ReferenceInput source="attrbutekval_id" reference="attrbuteval" filter={{ "key_id": form.getValues("attrbutekey_id") }}>
//                 <SelectInput
//                     label="属性值"
//                     optionText={"value"}
//                     create={<CreateDialog children={<CreateCategoriesValFrom />} />}
//                 />
//             </ReferenceInput>
//         </>
//     )
// }

// export const CreateCategoriesValFrom = (props: any) => {
//     const {onCancel } = useCreateSuggestionContext(); // 获取对话框上下文
//     const form = useFormContext()
//     const [create] = useCreate();
//     const notify = useNotify();
//     const [value, setValue] = React.useState('');
//     const { attrbutekey_id } = props

//     const handleSubmit = (event: any) => {
//         event.preventDefault();
//         let keyId = attrbutekey_id || form.getValues('attrbutekey_id'); // 直接读取字段值
//         const data = {
//             keyId: keyId,
//             value: value,
//         }


//         create('attrbuteval', { data }, {
//             onSuccess: (newData) => {
//                 notify('保存成功', { type: 'success' });
//                 props.onSuccess?.(newData);
//                 onCancel();
//             }
//         });
//     };
//     return (

//         <form onSubmit={handleSubmit}>
//             <DialogContent>

//                 <MuiTextField
//                     fullWidth
//                     margin="normal"
//                     label="请输入属性值"
//                     value={value}
//                     onChange={(event: any) => setValue(event.target.value)}
//                     autoFocus
//                 />
//             </DialogContent>

//             <DialogActions>
//                 <Button type="submit">提交</Button>
//                 <Button onClick={onCancel}>取消</Button>
//             </DialogActions>
//         </form>

//     )
// }
// export const CreateCategoriesKeyFrom = () => {
//     return (
//         <>
//             <Create redirect={false}>
//                 <SimpleForm>
//                     <TextInput source="name" label="名称" />
//                 </SimpleForm>
//             </Create>

//         </>
//     )
// }
// //导出此组件

// export const CreateDialog = (props: any) => {
//     const { filter, onCancel, onCreate } = useCreateSuggestionContext();

//     return (
//         <Dialog open onClose={onCancel}>

//             {React.cloneElement(props.children, {
//                 // 将对话框操作方法传递给子组件
//                 onCreate: onCreate,
//                 onCancel: onCancel
//             })}

//         </Dialog>
//     );
// };

// src/views/attributes/AttributeForm.tsx
import { Create } from "react-admin";
import AttributeForm from "./AttributeForm";
const AttributeCreate = () => (
    <Create redirect="list">
        <AttributeForm />
    </Create>
);
export default AttributeCreate;

