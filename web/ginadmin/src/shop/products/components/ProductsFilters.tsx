import React from "react";
import { useController } from "react-hook-form";
import { SearchInput, SelectInput, ReferenceInput } from "react-admin";
import { ToggleButton, ToggleButtonGroup } from "@mui/material";

interface ProductsFilterButtonsProps {
    // 这里可以根据实际情况添加更多的 props 类型定义
    alwaysOn?: boolean;
    name?: string; // 添加 name 属性以便于 useController 正确识别
}

const ProductsFilterButtons: React.FC<ProductsFilterButtonsProps> = (props) => {
    const {
        field,
        fieldState: { invalid, error }
    } = useController({ name: "status", defaultValue: 'all' });

    const handleAlignment = (event: React.SyntheticEvent, newAlignment: string) => {
        field.onChange(newAlignment);
    };

    return (
        <ToggleButtonGroup
            {...field}
            color="success"
            value={field.value || "all"} // 默认值为 "all"
            exclusive
            onChange={handleAlignment}
        >
            <ToggleButton key="all" value="all" aria-label="全部商品">
                全部商品
            </ToggleButton>
            <ToggleButton key="active" value="active" aria-label="发布中的">
                发布中的商品
            </ToggleButton>
            <ToggleButton key="pulled" value="pulled" aria-label="下架上">
                下架商品
            </ToggleButton>
            <ToggleButton key="disabled" value="disabled" aria-label="草稿的">
                草稿商品
            </ToggleButton>
        </ToggleButtonGroup>
    );
};




export const ProductsFilters = [
    <ProductsFilterButtons key={1} alwaysOn />,

    <SearchInput source="q" key={2} alwaysOn />,
    // 进行分类过滤
    <ReferenceInput key={3} source="product_category_id" label="分类过滤" reference="categories">
        <SelectInput optionText="name" label="分类过滤" />
    </ReferenceInput>,
    // 进行品牌过滤
    <ReferenceInput key={4} source="brand_id" label="品牌过滤" reference="brand">
        <SelectInput optionText="name" label="品牌过滤" />
    </ReferenceInput>
    // 其他过滤器可以在这里添加
];


