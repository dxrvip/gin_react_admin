import { Stack } from '@mui/material';
import {
    List,
    Datagrid,
    TextField,
    NumberField,
    DateField,
    ReferenceField,
    EditButton,
    ShowButton,
    DeleteButton,

} from 'react-admin';
import StatusField from '../../components/StatusField';
import { SecondHandProductList } from './SecondHandProduct/SecondHandProductList';
import { SecondHandProductCreate } from './SecondHandProduct/SecondHandProductCreate';
import { ProductsFilters } from './components/ProductsFilters';




export const ProductList = () => {

    return (
        <List title="商品列表"
            filters={ProductsFilters}
        >
            <Datagrid
                rowClick={() => {
                    return false; // 阻止默认行点击行为
                }}
            > 
                <TextField source="id" />
                <TextField source="title" label="机器型号" />
                <NumberField source="price" options={{ style: 'currency', currency: 'CNY' }} label="价格" />
                <NumberField source="stock" label="库存" />
                <ReferenceField source="productCategoryID" reference="categories" label="分类">
                    <TextField source="name" />
                </ReferenceField>
                <ReferenceField source="brandID" reference="brand" label="品牌">
                    <TextField source="name" label="品牌" />
                </ReferenceField>
                <DateField source="createdAt" showTime label="添加时间" />
                <StatusField source="status" label="状态" />
                <Stack direction="row" spacing={0}>
                    <EditButton key="edit-button" />
                   
                    {/* <ShowButton /> */}
                    <DeleteButton key="delete-button" />
                    <SecondHandProductList />
                </Stack>
            </Datagrid>
        </List>
    )
}

