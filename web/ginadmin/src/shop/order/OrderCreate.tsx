import {
    Create,
    SimpleForm,
    TextInput,
    NumberInput,
    SelectInput,
    required,
    ReferenceInput,
    ArrayInput,
    SimpleFormIterator,
    ButtonProps,
    useSimpleFormIterator,
    minLength,
    DateTimeInput,
} from 'react-admin';
import Grid from "@mui/material/Grid2"
import { Stack, Button, Box } from '@mui/material';
import SelectProductField from './components/SelectProductField';
import { OrderForm } from './components/OrderForm';


export const orderStatusChoices = [
    { id: 'pending', name: '待付款' },
    { id: 'paid', name: '已付款' },
    { id: 'shipping', name: '已发货' },
    { id: 'completed', name: '已完成' },
    { id: 'cancelled', name: '已取消' },
];

export const AddOrderButton = (props: ButtonProps) => {
    const { add } = useSimpleFormIterator();
    // const translate = useTranslate();

    return (
        <Box sx={{ m: "10px" }}>

            <Button variant="contained" color="success" onClick={() => add()}>添加商品</Button>


        </Box>
    );
};
export const RemoveOrderButton = (props: ButtonProps) => {
    const { remove } = useSimpleFormIterator();
    // const translate = useTranslate();

    return (
        <Box sx={{ m: "10px" }}>

            <Button variant="contained" color="warning" onClick={() => remove(0)}>清空商品</Button>



        </Box>
    );
};




function OrderCreate() {
    return (
        <Create title="添加订单" redirect="list">
            <SimpleForm defaultValues={{
                status: 'paid',
            }}>
                    <OrderForm />
            </SimpleForm>
        </Create>
    );
}

export default OrderCreate;