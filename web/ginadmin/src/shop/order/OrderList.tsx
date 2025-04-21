import { Stack } from '@mui/material';
import {
    List,
    Datagrid,
    TextField,
    DateField,
    NumberField,
    TextInput,
    SelectInput,
    FilterButton,
    CreateButton,
    TopToolbar,
    ExportButton,
    EditButton,
    ShowButton,
    ReferenceField,

    ReferenceInput,
} from 'react-admin';

export const ORDER_STATUS_CHOICES = [
    { id: 'pending', name: '待付款' },
    { id: 'paid', name: '已付款' },
    { id: 'shipping', name: '已发货' },
    { id: 'completed', name: '已完成' },
    { id: 'cancelled', name: '已取消' },
];

const OrderFilters = [
    <TextInput source="orderNo" label="订单号" alwaysOn />,
    <SelectInput source="status" label="订单状态" choices={ORDER_STATUS_CHOICES} />,
    <ReferenceInput source="user_id" reference="users" label="用户">
        <SelectInput optionText="username" />
    </ReferenceInput>
];

const ListActions = () => (
    <TopToolbar>
        <FilterButton />
        <CreateButton />
        <ExportButton />
    </TopToolbar>
);

function OrderList() {
    return (
        <List
            filters={OrderFilters}
            actions={<ListActions />}
            sort={{ field: 'created_at', order: 'DESC' }}
        >
            <Datagrid>
                <TextField source="orderNo" label="订单号" />
                <ReferenceField source="userId" reference="user" label="用户">
                    <TextField source="username" />
                </ReferenceField>
                <TextField source="address" label="收货地址" />
                {/* <ReferenceField source="address_id" reference="addresses" label="收货地址">
                </ReferenceField> */}
                <NumberField source="totalAmount" label="订单金额" options={{ style: 'currency', currency: 'CNY' }} />
                <TextField source="status" label="订单状态" />
                {/* <ArrayField source="orderItems" label="订单商品">
                    <SingleFieldList>
                        <ChipField source="product.title" />
                    </SingleFieldList>
                </ArrayField> */}
                <TextField source="note" label="备注" />
                {/* <DateField source="payment_time" label="支付时间" showTime />
                <DateField source="shipping_time" label="发货时间" showTime />
                <DateField source="completed_time" label="完成时间" showTime /> */}
                <DateField source="createAd" label="创建时间" showTime />
                <Stack direction="row" spacing={0}>
                    <EditButton />
                    <ShowButton />
                </Stack>
            </Datagrid>
        </List>
    );
}

export default OrderList;