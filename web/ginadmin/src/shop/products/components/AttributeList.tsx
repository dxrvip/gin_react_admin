import { useFormContext } from 'react-hook-form';
import {
    useDataProvider,

} from 'react-admin';
import { Attribute } from '../ProductsCreate';
import { Alert, Box, CircularProgress, Stack } from '@mui/material';
import { AttributeInput } from './AttributeInput';
import { useEffect, useState } from 'react';

export const AttributeList = () => {
    const form = useFormContext();
    const dataProvider = useDataProvider();
    const [attributes, setAttributes] = useState<Attribute[]>();
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState();


    const categoryID = form.getValues("productCategoryID");

    // const [values, setValues] = useState([]); // 初始化为一个空数组，用于存储属性值
    useEffect(() => {
        if (categoryID) {
            // 当 categoryID 变化时，重新获取数据
            dataProvider.getList<Attribute>('attribute', {
                pagination: { page: 1, perPage: 100 },
                sort: { field: "id", order: "ASC" },
                filter: { categories_id: categoryID },
            }).then(({ data }) => {
                setAttributes(data);
                setLoading(false);
            }).catch((err) => {
                setError(err);
                setLoading(false);
            })

        }
    }, [categoryID]);
    
 

    if (!categoryID) return <Alert severity="info">请先选择商品分类</Alert>;
    if (error) return <Alert severity="error">加载属性失败: {error}</Alert>;
    if (loading) return <CircularProgress />;

    return (
        <Box sx={{ width: '100%' }}>
            <Stack direction={"row"} spacing={1} sx={{"flexWrap": "wrap"}}>
                {/* 分配数组给子组建存选中value */}
                {attributes?.map((attr: any) => (
                    <Box key={attr.id} sx={{ minWidth: "70px" }}>
                        <AttributeInput
                            attribute={{
                                ...attr,    // 其他属性
                                options: attr.options || [],
                                min: attr.MinValue ?? undefined,
                                max: attr.MaxValue ?? undefined,
                            }}
                        />
                    </Box>
                ))}
            </Stack>
        </Box>
    );
};