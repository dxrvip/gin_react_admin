import {
  Edit,

  SimpleForm,

} from 'react-admin';
import ProductsForm from './components/ProductsForm';



export const ProductEdit = () => (
  <Edit>
    <SimpleForm>
      <ProductsForm />  
    </SimpleForm>
  </Edit>
);


export default ProductEdit;