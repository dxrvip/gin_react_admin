// src/utils/transformAttributes.ts
export interface AttributeData {
    attributeId: number;
    value: string | number;
}

export const transformAttributes = (data: any) => {
    // 提取动态属性字段
    console.log('transformAttributes', data);
    const attributes = Object.keys(data)
        .filter(key => key.startsWith('_attributes.'))
        .reduce((acc, key) => {
            const attrId = key.split('.')[1]; // 提取属性 ID（如 "1"）
            acc[attrId] = data[key];
            return acc;
        }, {} as Record<string, any>);

    // 2. 移除临时字段 _attributes
    const { _attributes, ...restData } = data;

    // 3. 返回最终数据
    return {
        ...restData,
        attributes
    };
};
