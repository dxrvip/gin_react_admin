import {
    SimpleForm,
    TextInput,
    SelectInput,
    ImageInput,
    required,
    minValue,
    maxValue,
    ImageField,
    BooleanInput,
    NumberInput,
    ReferenceInput,
    FormDataConsumer,
    DateInput,
    useCreate,
    useNotify,
    useRedirect,
} from "react-admin";
import { Box, Stack } from "@mui/material";
import { useState } from "react";
import { CONDITION_CHOICES, FUNCTION_CHOICES, SECONDHAND_STATUS, USAGE_DURATION_CHOICES } from "./SecondHandProductCreate";

const RegisterCode = () => {

    const now = new Date();
    const day = String(now.getDate()).padStart(2, '0');
    const hour = String(now.getHours()).padStart(2, '0');
    const minute = String(now.getMinutes()).padStart(2, '0');
    const second = String(now.getSeconds()).padStart(2, '0');
    const productCode = `${day}${hour}${minute}${second}`;
    return productCode
}

export const FormSeconHand = ({ record }: any) => {
    const [create] = useCreate()

    const { title, id, handleClose } = record

    // 获取通知方法
    const notify = useNotify();
    // 获取重定向方法
    const redirect = useRedirect();
    const postSave = (data: any) => {
        create("secondHandSkus", { data }, {
            onSuccess: (data) => onSuccess(data)
        })
    }
    function onSuccess(data: any) {
        notify(`成功创建二手商品：${title}`, { type: 'success' });
        setTimeout(handleClose, 1000); // 1秒后关闭
        // 刷新列表
        redirect("list")

    }
    return (
        <Box>
            <SimpleForm
                defaultValues={{
                    productsType: () => {
                        return `${title}-${RegisterCode()}`
                    },
                    productId: id || undefined,
                    function: 'perfect',
                    usageDuration: 'three_years',
                    freeShipping: true,
                    status: 'active',
                    isRepair: false,
                    accessories: true,
                    stock: 1,
                    batteryLife: undefined,
                    condition: 'obvious',
                }}
                shouldUnregister
                onSubmit={postSave}>
                {/* 价格 */}
                <TextInput source='title' label="商品标题" validate={[maxValue(60)]} />

                <ReferenceInput source='productId' reference="product">
                    <SelectInput label="选择父商品" />
                </ReferenceInput>
                <Stack direction="row" gap={2}>
                    <SelectInput source="condition" label="成色" choices={CONDITION_CHOICES} validate={[required()]} />
                    <SelectInput source="function" label="功能状态" choices={FUNCTION_CHOICES} validate={[required()]} />
                    <SelectInput source="usageDuration" label="使用年限" choices={USAGE_DURATION_CHOICES} validate={[required()]} />
                    <NumberInput source="price" label="价格"
                        helperText="代发拿货价格"
                        validate={[required(), minValue(0), maxValue(1000000)]} />
                    <NumberInput
                        source="stock"
                        label="数量（整数）"
                        validate={[required(), minValue(0), maxValue(1000000)]}
                    />
                </Stack>
                <Stack direction="row" spacing={1}>
                    <TextInput source="cost" label="成本价格" helperText="成本加上损耗" />
                    <TextInput source="productsType" label="货号" />
                    {/* 添加电池工作时间 */}
                    <NumberInput source="batteryLife" label="电池续航时间" helperText="分钟" />
                    {/* 机器工作时间（分钟） */}
                    <NumberInput source="workingTime" label="机器工作时间" helperText="分钟" />
                </Stack>
                <Stack direction="row" sx={{ width: "100%" }} spacing={1} >
                    {/* 配件是否齐全 */}
                    <BooleanInput
                        source="accessories"
                        label="配件是否齐全"
                        helperText="选择否后，需要填写缺少配件信息"
                        validate={[required()]} />

                    {/* <TextInput source="accessories" label="缺少信息" validate={[required()]} /> */}
                    {/* 是否包邮 */}
                    <BooleanInput source="freeShipping" label="是否包邮" validate={[required()]} />
                    {/* 运费 */}
                    <FormDataConsumer>
                        {
                            ({ formData }) =>
                                formData.freeShipping == false ? <NumberInput
                                    sx={{ width: "150px" }}
                                    variant="standard"
                                    source="shippingFee"
                                    label="运费金额"
                                    validate={[required()]} /> : null
                        }
                    </FormDataConsumer>
                    <BooleanInput source="isRepair" label="是否再保修" validate={[required()]} />
                    <FormDataConsumer>
                        {({ formData }) =>
                            formData.isRepair === true ? <DateInput
                                source="repairEndDate"
                                label="保修截止日期"
                                sx={{ width: "200px" }}
                                validate={[required()]} /> : null
                        }
                    </FormDataConsumer>
                    {/* <TextInput source="shipp  ingFee" label="运费" defaultValue={0} validate={[required(), minValue(0), maxValue(100)]} /> */}
                </Stack>
                <FormDataConsumer>
                    {
                        ({ formData }) =>
                            formData.accessories === false ? <TextInput
                                source="accessoriesList"
                                label="缺少配件信息"
                                validate={[required()]} /> : null
                    }
                </FormDataConsumer>
                {/* 图片 */}
                {/* <AccessoriesInput source="accessoriesList" /> */}
                {/* 增加换行 */}

                <Stack sx={{ mt: 2 }} direction={"row"} spacing={1}>
                    <ImageInput
                        source="picture"
                        label="商品图片"
                        validate={[required()]}
                        multiple
                        maxSize={500000000}>
                        <ImageField source="src" title="title" />
                    </ImageInput>

                </Stack>


                {/* 描述 */}
                <TextInput source="description" label="描述" />
                {/* 状态 */}
                <SelectInput source="status" choices={SECONDHAND_STATUS} label="状态" validate={required()} />

            </SimpleForm>
        </Box>
    )
}