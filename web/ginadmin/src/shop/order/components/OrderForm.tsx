import { Box, Stack } from "@mui/material"
import { ArrayInput, DateTimeInput, minValue, NumberInput, ReferenceInput, required, SelectInput, SimpleFormIterator, TextInput } from "react-admin"
import SelectProductField from "./SelectProductField"
import { AddOrderButton } from "../OrderCreate";
import { ORDER_STATUS_CHOICES } from "../OrderList";

const validateRequired = required();

export const OrderForm = () => {

    return <Box>

        <ArrayInput source="items" label="订单商品">
            <SimpleFormIterator fullWidth={false} inline
                getItemLabel={index => `#${index + 1}`}
                addButton={<AddOrderButton label={"添加商品"} />}
            // removeButton={<RemoveOrderButton label="清空商品" />}
            >
                <Stack direction="row" spacing={2} sx={{ alignContent: "flex-end" }}>
                    <ReferenceInput
                        source="productId"
                        reference="product"
                        label="商品"
                        filter={{ "is_second_hand_sku": true }} // 仅显示二手商品
                    >
                        <SelectProductField source="productId" />
                    </ReferenceInput>
                    <NumberInput source="quantity" label="数量" size="medium" validate={[validateRequired]} min={1} />
                    <NumberInput source="price" label="金额" size="medium" validate={[validateRequired]} min={0.01} step={0.01} />
                </Stack>
            </SimpleFormIterator>

        </ArrayInput>
        <Stack direction="row" spacing={2}>
            <DateTimeInput source="createdAt" label="订单时间" defaultValue={new Date()} validate={[validateRequired]} />
            <NumberInput source="costPrice" label="成本" validate={[validateRequired, minValue(0.01)]} />
            <NumberInput source="weight" label="重量(Kg)" validate={[validateRequired, minValue(0.01)]} />
            <ReferenceInput source="userId" reference="user" label="售卖用户">
                <SelectInput optionText="username" validate={[validateRequired]} />
            </ReferenceInput>
        </Stack>

        {/* <ReferenceInput source="address_id" reference="addresses" label="收货地址">
                    <SelectInput optionText="address" validate={[validateRequired]} />
                </ReferenceInput> */}


        <TextInput source="address" label="收货地址" validate={[validateRequired]} />
        <TextInput source="note" label="订单备注" multiline rows={3} />
        {/* 订单状态 */}
        <SelectInput source="status" label="订单状态" choices={ORDER_STATUS_CHOICES} validate={[required()]} />

    </Box>
}