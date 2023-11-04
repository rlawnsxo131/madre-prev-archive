import type { Meta, StoryObj } from '@storybook/react';

import { FlexLayout } from '../_layout/FlexLayout';
import type { ButtonProps } from './Button';
import { Button } from './Button';

const meta = {
  title: 'ui/Button',
  component: Button,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Button>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: 'button',
    theme: 'primary',
    size: 'medium',
    radius: 'medium',
    fullWidth: false,
    disabled: false,
  },
  render: (args) => {
    const itemBoxStyle = {
      width: '100%',
      marginTop: '1rem',
    };

    const _args: ButtonProps<'button'> = {
      children: args.children,
      radius: args.radius,
      style: { marginRight: '1rem' },
    };

    return (
      <FlexLayout flexDirection="column" alignItems="center">
        <FlexLayout
          justifyContent="center"
          style={{ width: itemBoxStyle.width }}
        >
          <Button {...args} as="a" />
        </FlexLayout>

        {/* variant full */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="primary" size="small" {..._args} />
          <Button theme="primary" {..._args} />
          <Button theme="primary" size="large" {..._args} />
          <Button theme="primary" size="large" fullWidth {..._args} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="secondary" size="small" {..._args} />
          <Button theme="secondary" {..._args} />
          <Button theme="secondary" size="large" {..._args} />
          <Button theme="secondary" size="large" fullWidth {..._args} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" size="small" {..._args} />
          <Button theme="warn" {..._args} />
          <Button theme="warn" size="large" {..._args} />
          <Button theme="warn" size="large" fullWidth {..._args} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button size="small" disabled {..._args} />
          <Button disabled {..._args} />
          <Button size="large" disabled {..._args} />
          <Button size="large" disabled fullWidth {..._args} />
        </FlexLayout>

        {/* variant outline */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="primary" variant="outline" size="small" {..._args} />
          <Button theme="primary" variant="outline" {..._args} />
          <Button theme="primary" variant="outline" size="large" {..._args} />
          <Button
            theme="primary"
            variant="outline"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="secondary" variant="outline" size="small" {..._args} />
          <Button theme="secondary" variant="outline" {..._args} />
          <Button theme="secondary" variant="outline" size="large" {..._args} />
          <Button
            theme="secondary"
            variant="outline"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" variant="outline" size="small" {..._args} />
          <Button theme="warn" variant="outline" {..._args} />
          <Button theme="warn" variant="outline" size="large" {..._args} />
          <Button
            theme="warn"
            variant="outline"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button variant="outline" size="small" disabled {..._args} />
          <Button variant="outline" disabled {..._args} />
          <Button variant="outline" size="large" disabled {..._args} />
          <Button
            variant="outline"
            size="large"
            fullWidth
            disabled
            {..._args}
          />
        </FlexLayout>

        {/* variant ghost */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="primary" variant="ghost" size="small" {..._args} />
          <Button theme="primary" variant="ghost" {..._args} />
          <Button theme="primary" variant="ghost" size="large" {..._args} />
          <Button
            theme="primary"
            variant="ghost"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="secondary" variant="ghost" size="small" {..._args} />
          <Button theme="secondary" variant="ghost" {..._args} />
          <Button theme="secondary" variant="ghost" size="large" {..._args} />
          <Button
            theme="secondary"
            variant="ghost"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" variant="ghost" size="small" {..._args} />
          <Button theme="warn" variant="ghost" {..._args} />
          <Button theme="warn" variant="ghost" size="large" {..._args} />
          <Button
            theme="warn"
            variant="ghost"
            size="large"
            fullWidth
            {..._args}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button variant="ghost" size="small" disabled {..._args} />
          <Button variant="ghost" disabled {..._args} />
          <Button variant="ghost" size="large" disabled {..._args} />
          <Button variant="ghost" size="large" fullWidth disabled {..._args} />
        </FlexLayout>
      </FlexLayout>
    );
  },
};
