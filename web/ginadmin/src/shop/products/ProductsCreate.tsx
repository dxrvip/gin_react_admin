
import {
    Create,
    SimpleForm
} from 'react-admin';

import ProductsForm from './components/ProductsForm';
export const STSATUS = [
    { id: 'active', name: '发布' },
    { id: 'disabled', name: '草稿' },
    { id: 'pulled', name: '下架' }
]
export const ProductCreate = () => {

    // 处理属性
    return <Create
        title="新建商品"
        redirect="list"
        >
        <SimpleForm sx={{ width: "80%" }} shouldUnregister>
           <ProductsForm />
        </SimpleForm>
    </Create >
};

export interface Attribute {
    id: number;
    name: string;
    isRequired: boolean;
    defaultValue: string;
    options: Options[];
}
export interface Options {
    label: string;
    value: string;
}
