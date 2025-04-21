// src/attributes/AttributeForm.tsx
import {
  SimpleForm,
  TextInput,
  NumberInput,
  SelectInput,
  BooleanInput,
  ArrayInput,
  SimpleFormIterator,
  required,
  FormDataConsumer,
  ReferenceArrayInput,
  SelectArrayInput,
  minValue,
  maxValue
} from "react-admin";
import { Stack } from "@mui/material";
const AttributeForm = (props: any) => {
  return (
    <SimpleForm
      // 初始化默认值
      defaultValues={{
        options: [],
        categoryIds: [],
        isRequired: false,
        min: 0,
        max: 100
      }}
    >
      {/* 公共字段 */}
      <Stack direction={"row"} spacing={3}>
        <TextInput source="name" label="属性名称" validate={[required()]} fullWidth />

        <SelectInput
          source="type"
          label="数据类型"
          choices={[
            { id: "string", name: "文本" },
            { id: "number", name: "数字" },
            { id: "enum", name: "枚举值" }
          ]}
          validate={[required()]}
        />
        <TextInput
          source="unit"
          label="单位"
          helperText="例如: 元、kg、个等"
          fullWidth
        />
      </Stack>

      {/* 动态字段区域 */}
      <FormDataConsumer>
        {({ formData }) => {
          switch (formData.type) {
            case "number":
              return (
                <Stack direction={"row"} spacing={3}>
                  <NumberInput
                    source="min"
                    label="最小值"
                    validate={[minValue(0)]}
                    sx={{ width: '120px' }}
                  />
                  <NumberInput
                    source="max"
                    label="最大值"
                    validate={[maxValue(10000)]}
                    sx={{ width: '120px' }}
                  />

                </Stack>
              );
            case "enum":
              return (
                <ArrayInput
                  source="options"
                  label="枚举选项"
                  validate={[required()]}
                  helperText="至少需要定义一个选项"
                >
                  <SimpleFormIterator inline>
                    <Stack direction={"row"} spacing={3}>
                      <TextInput source="value" label="选项值" validate={[required()]} />
                      <TextInput source="label" label="显示名称" />
                    </Stack>
                  </SimpleFormIterator>
                </ArrayInput>
              );

            default:
              return null;
          }
        }}
      </FormDataConsumer>

      <Stack direction={"row"} spacing={2}>
        <TextInput
          source="defaultValue"
          label="默认数值"
          helperText="用户未输入时将使用此默认值"
        />
        <BooleanInput sx={{width: "200px"}} source="isRequired" label="是否必填"  fullWidth/>

      </Stack>


      {/* 分类关联 */}
      <ReferenceArrayInput
        source="categoryIds"
        reference="categories"
        label="适用分类"
        parse={(ids) => ids?.map((id: number) => ({ id })) || []}
        format={(categories) => categories?.map((c: any) => c.id) || []}
      >
        <SelectArrayInput optionText="name" />
      </ReferenceArrayInput>
    </SimpleForm>
  );
};

export default AttributeForm;