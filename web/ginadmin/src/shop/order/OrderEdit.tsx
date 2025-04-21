import {
    Edit,
    SimpleForm,
} from 'react-admin';
import { OrderForm } from './components/OrderForm';



function OrderEdit() {
    return (
        <Edit>
            <SimpleForm>
                 <OrderForm />
            </SimpleForm>
        </Edit>
    );
}

export default OrderEdit;