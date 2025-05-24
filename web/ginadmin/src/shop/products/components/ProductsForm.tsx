import { Box, Stack } from "@mui/material";
import { FormDataConsumer, ImageField, ImageInput, maxLength, minLength, NumberInput, ReferenceInput, required, SelectInput, TextInput } from "react-admin";
import { AttributeList } from "./AttributeList";
import { STSATUS } from "../ProductsCreate";



export default function ProductsForm() {
    return (
        <Box>

            <TextInput source="title" label="商品标题" validate={[required(), maxLength(255), minLength(2)]} />

            <Stack direction="row" spacing={1}>

                <ReferenceInput source="productCategoryID" reference="categories">
                    <SelectInput optionText="name" validate={[required()]} label="商品分类" />
                </ReferenceInput>
                <ReferenceInput source="brandID" reference="brand">
                    <SelectInput
                        optionText="name"
                        optionValue="id"   
                        validate={[required()]}
                        label="品牌选择"
                        parse={v => Number(v)}  // 添加类型转换
                    />
                </ReferenceInput>


                <NumberInput source="price" validate={[required()]} min={0} step={0.01} label="商品价格" />
                <NumberInput source="stock" validate={[required()]} min={0} label="库存数量" />
            </Stack>
            <FormDataConsumer>
                {({ formData }) => (
                    formData?.productCategoryID && <AttributeList />
                )}
            </FormDataConsumer>



            <ImageInput source="image" label="选择图片" accept={{ 'image/*': ['.png', '.jpg', '.jpeg'] }}>
                <ImageField source="src" title="title" />
            </ImageInput>

            <TextInput source="description" multiline fullWidth label="描述" />
            <SelectInput source="status" choices={STSATUS} defaultValue="active" label="状态"
                validate={required()}
            />
        </Box>

    ); // 这个组件本身不需要渲染任何内容
}