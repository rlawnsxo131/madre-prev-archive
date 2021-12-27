import {
  Column,
  CreateDateColumn,
  Entity,
  Index,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';
import { User } from '../../user';

@Entity('data')
export class Data {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Column({ length: 36 })
  user_id!: string;

  @Column({ length: 255 })
  file_url!: string;

  @Column()
  title!: string;

  @Column({ nullable: true, default: null })
  description?: string;

  @Column({ default: false })
  is_public!: boolean;

  @CreateDateColumn({ type: 'timestamp' })
  created_at!: Date;

  @UpdateDateColumn({ type: 'timestamp' })
  updated_at!: Date;

  @ManyToOne(() => User, (user) => user.datas, {
    createForeignKeyConstraints: false,
  })
  @Index('ix_user_id')
  @JoinColumn({ name: 'user_id', referencedColumnName: 'id' })
  user!: User;
}
