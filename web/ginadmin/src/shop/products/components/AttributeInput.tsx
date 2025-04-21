import {
    required,
    minValue,
    maxValue,
    SelectInput,
    TextInput,
    NumberInput,
    maxLength,
    minLength,
} from 'react-admin';
import { Options } from '../ProductsCreate';

export interface AttributeInputProps {
    attribute: {
        id: number;
        name: string;
        type: 'enum' | 'number' | 'string';
        defaultValue: string;
        isRequired: boolean;
        options: Options[];
        min?: number;
        max?: number;
        value?: any; // 用于编辑时的当前值
    };
}

export const AttributeInput = ({ attribute }: AttributeInputProps) => {
    const { id, type, name, options, isRequired, min, max, defaultValue } = attribute;
    const fieldName = `attributes._${id}`; // ✅ 字段名格式：attributes.1、attributes.2
    const validate: any[] = isRequired ? [required()] : []; // 如果是必填项，则添加 required 验证


    switch (type) {
        case 'enum':
            return (
 
                <SelectInput source={fieldName} // 确保 source 格式正确
                    label={name}
                    choices={options.map((o: Options) => ({
                        id: o.value, // 使用 value 作为 id
                        name: o.label, // 显示名称
                    }))}
                    validate={isRequired ? [required()] : undefined} />
            );

        case 'number':
            return (
                <NumberInput
                    source={fieldName} // 确保 source 格式正确
                    type="number"
                    label={name}
                    defaultValue={defaultValue ?? ""} // 设置默认值
                    validate={[...validate, minValue(min ?? 0), maxValue(max ?? 50000)]}
                />
            );

        case 'string':
            return (
                <TextInput
                    source={fieldName} // 确保 source 格式正确
                    type="text"
                    label={name}
                    defaultValue={defaultValue ?? ""} // 设置默认值
                    validate={[...validate, maxLength(40), minLength(1)]} // 添加最大长度和最小长度验证
                />
            );

        default:
            return null;
    }
};