import type { Meta, StoryObj } from '@storybook/react';

import { FlexLayout } from '../FlexLayout';
import type { ButtonProps } from './Button';
import { Button } from './Button';

const meta = {
  title: 'atoms/Button',
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
      width: 'calc(100% - 2rem)',
      marginTop: '1rem',
    };

    const argsStyle: ButtonProps<'button'> = {
      children: args.children,
      radius: args.radius,
    };

    return (
      <FlexLayout
        flex="1"
        flexDirection="column"
        alignItems="center"
        style={{ width: '100%', height: '100%', padding: '1rem 0' }}
      >
        <FlexLayout
          justifyContent="center"
          style={{ width: itemBoxStyle.width }}
        >
          <Button {...args} />
        </FlexLayout>

        {/* variant full */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="primary" size="small" {...argsStyle} />
          <Button theme="primary" {...argsStyle} />
          <Button theme="primary" size="large" {...argsStyle} />
          <Button theme="primary" size="large" fullWidth {...argsStyle} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="secondary" size="small" {...argsStyle} />
          <Button theme="secondary" {...argsStyle} />
          <Button theme="secondary" size="large" {...argsStyle} />
          <Button theme="secondary" size="large" fullWidth {...argsStyle} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" size="small" {...argsStyle} />
          <Button theme="warn" {...argsStyle} />
          <Button theme="warn" size="large" {...argsStyle} />
          <Button theme="warn" size="large" fullWidth {...argsStyle} />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button size="small" disabled {...argsStyle} />
          <Button disabled {...argsStyle} />
          <Button size="large" disabled {...argsStyle} />
          <Button size="large" disabled fullWidth {...argsStyle} />
        </FlexLayout>

        {/* variant outline */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button
            theme="primary"
            variant="outline"
            size="small"
            {...argsStyle}
          />
          <Button theme="primary" variant="outline" {...argsStyle} />
          <Button
            theme="primary"
            variant="outline"
            size="large"
            {...argsStyle}
          />
          <Button
            theme="primary"
            variant="outline"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button
            theme="secondary"
            variant="outline"
            size="small"
            {...argsStyle}
          />
          <Button theme="secondary" variant="outline" {...argsStyle} />
          <Button
            theme="secondary"
            variant="outline"
            size="large"
            {...argsStyle}
          />
          <Button
            theme="secondary"
            variant="outline"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" variant="outline" size="small" {...argsStyle} />
          <Button theme="warn" variant="outline" {...argsStyle} />
          <Button theme="warn" variant="outline" size="large" {...argsStyle} />
          <Button
            theme="warn"
            variant="outline"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button variant="outline" size="small" disabled {...argsStyle} />
          <Button variant="outline" disabled {...argsStyle} />
          <Button variant="outline" size="large" disabled {...argsStyle} />
          <Button
            variant="outline"
            size="large"
            fullWidth
            disabled
            {...argsStyle}
          />
        </FlexLayout>

        {/* variant ghost */}
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="primary" variant="ghost" size="small" {...argsStyle} />
          <Button theme="primary" variant="ghost" {...argsStyle} />
          <Button theme="primary" variant="ghost" size="large" {...argsStyle} />
          <Button
            theme="primary"
            variant="ghost"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button
            theme="secondary"
            variant="ghost"
            size="small"
            {...argsStyle}
          />
          <Button theme="secondary" variant="ghost" {...argsStyle} />
          <Button
            theme="secondary"
            variant="ghost"
            size="large"
            {...argsStyle}
          />
          <Button
            theme="secondary"
            variant="ghost"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button theme="warn" variant="ghost" size="small" {...argsStyle} />
          <Button theme="warn" variant="ghost" {...argsStyle} />
          <Button theme="warn" variant="ghost" size="large" {...argsStyle} />
          <Button
            theme="warn"
            variant="ghost"
            size="large"
            fullWidth
            {...argsStyle}
          />
        </FlexLayout>
        <FlexLayout alignItems="center" style={itemBoxStyle}>
          <Button variant="ghost" size="small" disabled {...argsStyle} />
          <Button variant="ghost" disabled {...argsStyle} />
          <Button variant="ghost" size="large" disabled {...argsStyle} />
          <Button
            variant="ghost"
            size="large"
            fullWidth
            disabled
            {...argsStyle}
          />
        </FlexLayout>
      </FlexLayout>
    );
  },
};
