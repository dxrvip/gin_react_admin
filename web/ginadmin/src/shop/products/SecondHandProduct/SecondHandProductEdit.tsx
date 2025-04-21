import React from 'react';
import {
  Edit,
  SimpleForm,
  ReferenceInput,
  SelectInput,
  NumberInput,
  TextInput,
  BooleanInput,
  ImageInput,
  ImageField,
  required,
  minValue,
} from 'react-admin';

const conditionChoices = [
  { id: 'new', name: '全新' },
  { id: 'like_new', name: '几乎全新' },
  { id: 'light', name: '轻微使用' },
  { id: 'obvious', name: '明显使用' },
  { id: 'serious', name: '重度使用' },
  { id: 'damaged', name: '有损坏' },
];

const functionChoices = [
  { id: 'perfect', name: '功能完好' },
  { id: 'repaired', name: '维修过' },
  { id: 'usable', name: '可用' },
  { id: 'unusable', name: '不可用' },
];

const usageDurationChoices = [
  { id: 'unused', name: '未使用' },
  { id: 'half_year', name: '半年内' },
  { id: 'one_year', name: '一年内' },
  { id: 'three_years', name: '三年内' },
];

export const SecondHandProductEdit = () => (
  <Edit>
    <SimpleForm>
      <ReferenceInput 
        source="productId" 
        reference="products" 
        validate={required() as any}
        label="关联商品"
      >
        <SelectInput optionText="name" />
      </ReferenceInput>
      
      <NumberInput 
        source="price" 
        validate={[required(), minValue(0)] as any}
        label="价格(元)"
      />
      
      <NumberInput 
        source="stock" 
        validate={[required(), minValue(0)] as any}
        label="库存"
      />
      
      <SelectInput 
        source="condition" 
        choices={conditionChoices}
        validate={required() as any}
        label="商品成色"
      />
      
      <SelectInput 
        source="function" 
        choices={functionChoices}
        validate={required() as any}
        label="功能状态"
      />
      
      <SelectInput 
        source="usageDuration" 
        choices={usageDurationChoices}
        validate={required() as any}
        label="使用时长"
      />
      
      <BooleanInput source="accessories" label="是否有配件" />
      <BooleanInput source="freeShipping" label="是否包邮" />
      
      <NumberInput 
        source="shippingFee"
        validate={minValue(0) as any}
        label="运费(元)"
      />
      
      <TextInput 
        source="description" 
        multiline
        rows={4}
        label="描述"
      />
      
      <ImageInput source="images" multiple label="商品图片">
        <ImageField source="url" title="title" />
      </ImageInput>
    </SimpleForm>
  </Edit>
);
