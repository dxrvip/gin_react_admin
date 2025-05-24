import {
    Edit,
    SimpleForm,
} from 'react-admin';
import { OrderForm } from './components/OrderForm';



function OrderEdit() {
    return (
        <Edit redirect="list">
            <SimpleForm>
                 <OrderForm />
            </SimpleForm>
        </Edit>
    );
}

export default OrderEdit;